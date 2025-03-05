import { StateGraph, MessagesAnnotation } from '@langchain/langgraph';
import fs from 'fs';
import { ChatOpenAI } from '@langchain/openai';
import { HumanMessage, SystemMessage } from '@langchain/core/messages';
import path from 'path';

require('dotenv').config({ path: './.env.local' });

// Define the agent state
interface AgentState {
  messages: string[]; // Tracks conversation history
  files?: string[]; // List of project files
  filesToFix?: string[]; // Files that need formatting
  fixedFiles?: { [key: string]: string }; // Stores formatted content
}

// Initialize OpenAI model
const model = new ChatOpenAI({
  modelName: 'gpt-4-turbo',
  temperature: 0,
});

// Create a stateful graph
const graph = new StateGraph(MessagesAnnotation);

// 1️⃣ Capture user message and store it in history
graph.addNode('get_user_prompt', async (state: AgentState, input: string) => ({
  ...state,
  messages: [...(state.messages || []), `User: ${input}`],
}));

// 2️⃣ List files in the project directory
graph.addNode('list_files', async (state: AgentState) => {
  const projectDir = './src'; // Change this directory as needed
  const files = fs.readdirSync(projectDir).filter((file) => file.endsWith('.txt'));
  return { ...state, files };
});

// 3️⃣ Ask LLM: Do we need to fix formatting?
graph.addNode('check_formatting', async (state: AgentState) => {
  if (!state.files || state.files.length === 0) return { ...state, filesToFix: [] };

  const messages = [
    new SystemMessage('You are an assistant that determines if files need formatting.'),
    new HumanMessage(
      `The user asked: "${
        state.messages[state.messages.length - 1]
      }". Here are the files:\n${state.files.join(
        '\n'
      )}\nWhich files need formatting? Respond with a JSON list of file names.`
    ),
  ];

  const response = await model.invoke(messages);

  try {
    const parsedResponse = JSON.parse(response.content);
    return {
      ...state,
      filesToFix: parsedResponse,
      messages: [...state.messages, `AI: Files to fix - ${parsedResponse.join(', ')}`],
    };
  } catch {
    return {
      ...state,
      filesToFix: state.files,
      messages: [...state.messages, `AI: Couldn't parse response, fixing all files.`],
    };
  }
});

// 4️⃣ Fix formatting of selected files
graph.addNode('fix_format', async (state: AgentState) => {
  if (!state.filesToFix || state.filesToFix.length === 0) return state;

  const fixedFiles: { [key: string]: string } = {};

  for (const file of state.filesToFix) {
    const filePath = path.join('./src', file);
    const content = fs.readFileSync(filePath, 'utf8');

    const messages = [
      new SystemMessage('Fix the formatting of this file.'),
      new HumanMessage(`File: ${file}\n\n${content}`),
    ];

    const response = await model.invoke(messages);
    fixedFiles[file] = response.content;
  }

  return { ...state, fixedFiles, messages: [...state.messages, `AI: Formatting completed.`] };
});

// 5️⃣ Save formatted files back to disk
graph.addNode('save_files', async (state: AgentState) => {
  if (!state.fixedFiles) return state;

  for (const [file, content] of Object.entries(state.fixedFiles)) {
    const filePath = path.join('./src', file);
    fs.writeFileSync(filePath, content, 'utf8');
  }

  console.log('✅ Files have been fixed and saved!');

  return { ...state, messages: [...state.messages, `AI: Saved formatted files.`] };
});

// Define execution flow
graph.setEntryPoint('get_user_prompt');
graph.addEdge('get_user_prompt', 'list_files');
graph.addEdge('list_files', 'check_formatting');
graph.addEdge('check_formatting', 'fix_format');
graph.addEdge('fix_format', 'save_files');

// Compile the agent
const agent = graph.compile();

// Run the agent
async function runAgent() {
  const userPrompt = 'Do I need to fix formatting in my project?';
  const state = await agent.invoke({ messages: [] }, userPrompt);

  console.log('🔹 Conversation History:\n', state.messages.join('\n'));
}

runAgent();

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
