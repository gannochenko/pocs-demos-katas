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
  // Define the schema for classification
  const classificationSchema = z.object({
    sentiment: z.enum(['happy', 'neutral', 'sad', 'aggressive']).describe('The sentiment of the text'),
    aggressiveness: z
      .number()
      .int()
      .min(1)
      .max(10)
      .describe('How aggressive the text is on a scale from 1 to 10'),
    language: z
      .enum(['spanish', 'english', 'french', 'german', 'italian', 'russian', 'other'])
      .describe('The language the text is written in'),
  });

  // Create the parser with your schema
  const parser = StructuredOutputParser.fromZodSchema(classificationSchema);

  // Update the prompt template to include the format instructions
  const taggingPrompt = ChatPromptTemplate.fromTemplate(
    `Extract the desired information from the following passage.
    
    {format_instructions}
    
    Passage:
    {input}
    `
  );

  // Create the chain
  const chain = taggingPrompt.pipe(model).pipe(parser);

  // Run the chain
  const result = await chain.invoke({
    input: 'Ублюдок, мать твою, а ну иди сюда говно собачье',
    format_instructions: parser.getFormatInstructions(),
  });

  console.log(result);
})().catch(console.error);
