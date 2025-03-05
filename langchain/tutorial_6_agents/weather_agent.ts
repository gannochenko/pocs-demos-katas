import { TavilySearchResults } from '@langchain/community/tools/tavily_search';
import { ChatOpenAI } from '@langchain/openai';
import { HumanMessage, AIMessage } from '@langchain/core/messages';
import { ToolNode } from '@langchain/langgraph/prebuilt';
import { StateGraph, MessagesAnnotation } from '@langchain/langgraph';

require('dotenv').config({ path: './.env.local' });

// Define the tools for the agent to use
const tools = [new TavilySearchResults({ maxResults: 3 })];
const toolNode = new ToolNode(tools);

// Create a model and give it access to the tools
const model = new ChatOpenAI({
  model: 'gpt-4o-mini',
  temperature: 0,
}).bindTools(tools);

// Define the function that determines whether to continue or not
function shouldContinue({ messages }: typeof MessagesAnnotation.State) {
  const lastMessage = messages[messages.length - 1] as AIMessage;

  // If the LLM makes a tool call, then we route to the "tools" node
  if (lastMessage.tool_calls?.length) {
    return 'tools';
  }
  // Otherwise, we stop (reply to the user) using the special "__end__" node
  return '__end__';
}

// Define the function that calls the model
async function callModel(state: typeof MessagesAnnotation.State) {
  const response = await model.invoke(state.messages);

  // We return a list, because this will get added to the existing list
  return { messages: [response] };
}

// Define a new graph
const workflow = new StateGraph(MessagesAnnotation)
  .addNode('agent', callModel)
  .addEdge('__start__', 'agent') // __start__ is a special name for the entrypoint
  .addNode('tools', toolNode)
  .addEdge('tools', 'agent')
  .addConditionalEdges('agent', shouldContinue);

// Finally, we compile it into a LangChain Runnable.
const app = workflow.compile();

(async function run() {
  // Use the agent
  let nextState = await app.invoke({
    messages: [new HumanMessage('what is the weather in sf')],
  });
  console.log(nextState.messages[nextState.messages.length - 1].content);

  nextState = await app.invoke({
    // Including the messages from the previous run gives the LLM context.
    // This way it knows we're asking about the weather in NY
    messages: [...nextState.messages, new HumanMessage('what about ny')],
  });
  console.log(nextState.messages[nextState.messages.length - 1].content);
})().catch(console.error);
