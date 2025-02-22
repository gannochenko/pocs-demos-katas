import { ChatOpenAI } from '@langchain/openai';
import { HumanMessage, SystemMessage } from '@langchain/core/messages';
import { ChatPromptTemplate } from '@langchain/core/prompts';

require('dotenv').config({ path: './.env.local' });

const model = new ChatOpenAI({ model: 'gpt-4o-mini-2024-07-18' });

(async function name() {
  //   const messages = [
  //     new SystemMessage('Translate the following from English into Italian'),
  //     new HumanMessage('hi!'),
  //   ];
  //   // one call
  //   await model.invoke(messages);
  //   // stream
  //   const stream = await model.stream(messages);
  //   const chunks = [];
  //   for await (const chunk of stream) {
  //     chunks.push(chunk);
  //     console.log(`${chunk.content}|`);
  //   }
  // prompt template:
  const systemTemplate = 'Translate the following from English into {language}';
  const promptTemplate = ChatPromptTemplate.fromMessages([
    ['system', systemTemplate],
    ['user', '{text}'],
  ]);

  const promptValue = await promptTemplate.invoke({
    language: 'italian',
    text: 'Give me more of those sweet, sweet prompts!',
  });

  const stream = await model.stream(promptValue);
  const chunks = [];

  for await (const chunk of stream) {
    chunks.push(chunk);
    console.log(`${chunk.content}|`);
  }
})().catch(console.error);

// import { ChatOpenAI } from "langchain/chat_models/openai";

// const llm = new ChatOpenAI();
// await llm.invoke("Hello, world!");
