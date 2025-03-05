import { TavilySearchResults } from '@langchain/community/tools/tavily_search';
import { ChatOpenAI } from '@langchain/openai';
import { MemorySaver } from '@langchain/langgraph';
import { HumanMessage } from '@langchain/core/messages';
import { createReactAgent } from '@langchain/langgraph/prebuilt';

require('dotenv').config({ path: './.env.local' });

const agentTools = [new TavilySearchResults({ maxResults: 3 })];
const agentModel = new ChatOpenAI({ temperature: 0 });

const agentCheckpointer = new MemorySaver();
const agent = createReactAgent({
  llm: agentModel,
  tools: agentTools,
  checkpointSaver: agentCheckpointer,
});

(async function run() {
  const agentFinalState = await agent.invoke(
    { messages: [new HumanMessage('what is the current weather in sf')] },
    { configurable: { thread_id: '42' } }
  );

  console.log(agentFinalState.messages[agentFinalState.messages.length - 1].content);
})().catch(console.error);
