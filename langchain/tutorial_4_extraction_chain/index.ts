import { ChatOpenAI, OpenAIEmbeddings } from '@langchain/openai';
import { Document } from '@langchain/core/documents';
import { PDFLoader } from '@langchain/community/document_loaders/fs/pdf';
import { RecursiveCharacterTextSplitter } from '@langchain/textsplitters';
import { QdrantVectorStore } from '@langchain/qdrant';
import { ChatPromptTemplate } from '@langchain/core/prompts';
import { z } from 'zod';
import { StructuredOutputParser } from '@langchain/core/output_parsers';

require('dotenv').config({ path: './.env.local' });

const model = new ChatOpenAI({
  model: 'gpt-4o-mini-2024-07-18',
  temperature: 0,
});

(async function run() {
  const personSchema = z.object({
    name: z.optional(z.string()).describe('The name of the person'),
    hair_color: z.optional(z.string()).describe("The color of the person's hair if known"),
    height_in_meters: z.number().nullish().describe('Height measured in meters'),
  });

  const dataSchema = z.object({
    people: z.array(personSchema).describe('Extracted data about people'),
  });

  const promptTemplate = ChatPromptTemplate.fromMessages([
    [
      'system',
      `You are an expert extraction algorithm.
  Only extract relevant information from the text.
  If you do not know the value of an attribute asked to extract,
  return null for the attribute's value.`,
    ],
    // Please see the how-to about improving performance with
    // reference examples.
    // ["placeholder", "{examples}"],
    ['human', '{text}'],
  ]);

  const structuredModel = model.withStructuredOutput(dataSchema, {
    name: 'person',
  });

  const prompt = await promptTemplate.invoke({
    text: 'My name is Jeff, my hair is black and i am 6 feet tall. Anna has the same color hair as me.',
  });
  const result = await structuredModel.invoke(prompt);

  console.log(result);
})().catch(console.error);
