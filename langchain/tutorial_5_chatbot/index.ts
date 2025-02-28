import { ChatOpenAI } from '@langchain/openai';
import { START, END, MessagesAnnotation, StateGraph, MemorySaver, Annotation } from '@langchain/langgraph';
import { v4 as uuidv4 } from 'uuid';
import { ChatPromptTemplate } from '@langchain/core/prompts';
import { SystemMessage, HumanMessage, AIMessage, trimMessages } from '@langchain/core/messages';

require('dotenv').config({ path: './.env.local' });

const llm = new ChatOpenAI({
  model: 'gpt-4o-mini-2024-07-18',
  temperature: 0,
});

(async function run() {
  const GraphAnnotation = Annotation.Root({
    ...MessagesAnnotation.spec,
    language: Annotation<string>(),
  });

  const trimmer = trimMessages({
    maxTokens: 10,
    strategy: 'last',
    tokenCounter: (msgs) => msgs.length,
    includeSystem: true,
    allowPartial: false,
    startOn: 'human',
  });

  const promptTemplate = ChatPromptTemplate.fromMessages([
    [
      'system',
      'You talk like a drunk Rick Sanchez. Answer all questions to the best of your ability in {language}.',
    ],
    ['placeholder', '{messages}'],
  ]);

  // Define the function that calls the model
  const callModel = async (state: typeof GraphAnnotation.State) => {
    const trimmedMessage = await trimmer.invoke(state.messages);
    const prompt = await promptTemplate.invoke({ ...state, messages: trimmedMessage });
    const response = await llm.invoke(prompt);
    return { messages: response };
  };

  // Define a new graph
  const workflow = new StateGraph(GraphAnnotation)
    // Define the node and edge
    .addNode('model', callModel)
    .addEdge(START, 'model')
    .addEdge('model', END);

  // Add memory
  const memory = new MemorySaver();
  const app = workflow.compile({ checkpointer: memory });

  const config = { configurable: { thread_id: uuidv4() } };

  const output = await app.invoke(
    {
      messages: [
        {
          role: 'user',
          content: 'Hi im bob',
        },
      ],
      language: 'Russian',
    },
    config
  );
  console.log(output.messages[output.messages.length - 1].content);
})().catch(console.error);
