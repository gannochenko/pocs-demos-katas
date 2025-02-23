import { ChatOpenAI, OpenAIEmbeddings } from '@langchain/openai';
import { Document } from '@langchain/core/documents';
import { PDFLoader } from '@langchain/community/document_loaders/fs/pdf';
import { RecursiveCharacterTextSplitter } from '@langchain/textsplitters';
import { QdrantVectorStore } from '@langchain/qdrant';

require('dotenv').config({ path: './.env.local' });

const model = new ChatOpenAI({ model: 'gpt-4o-mini-2024-07-18' });

(async function run() {
  // const documents = [
  //   new Document({
  //     pageContent: 'Dogs are great companions, known for their loyalty and friendliness.',
  //     metadata: { source: 'mammal-pets-doc' },
  //   }),
  //   new Document({
  //     pageContent: 'Cats are independent pets that often enjoy their own space.',
  //     metadata: { source: 'mammal-pets-doc' },
  //   }),
  // ];

  // console.log(documents);

  const loader = new PDFLoader('/Users/gannochenko/Downloads/cv.pdf');

  const docs = await loader.load();
  // console.log(docs[0].pageContent.slice(0, 200));
  // console.log(docs[0].metadata);

  const textSplitter = new RecursiveCharacterTextSplitter({
    chunkSize: 1000,
    chunkOverlap: 200,
  });

  const allSplits = await textSplitter.splitDocuments(docs);

  // console.log('Splits:', allSplits.length);
  // console.log(allSplits);

  const embeddings = new OpenAIEmbeddings({
    model: 'text-embedding-3-large',
  });

  const vector1 = await embeddings.embedQuery(allSplits[0].pageContent);
  const vector2 = await embeddings.embedQuery(allSplits[1].pageContent);

  const vectorStore = await QdrantVectorStore.fromExistingCollection(embeddings, {
    url: process.env.QDRANT_URL,
    collectionName: 'langchainjs-testing',
  });

  await vectorStore.addDocuments(allSplits);

  const results1 = await vectorStore.similaritySearch('Who is Sergei?');

  console.log(results1);
})().catch(console.error);
