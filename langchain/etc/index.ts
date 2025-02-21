import { ChatOpenAI } from "@langchain/openai";
import { initializeAgentExecutorWithOptions } from "langchain/agents";
import { DynamicTool } from "@langchain/core/tools";
import * as fs from "fs/promises";
import * as path from "path";

async function createFileOperationsAgent() {
  const model = new ChatOpenAI({
    temperature: 0,
    modelName: "gpt-4",
  });

  // Tool to create a new file
  const createFileTool = new DynamicTool({
    name: "create_file",
    description: "Creates a new file with the given name",
    func: async (fileName: string) => {
      try {
        // Ensure the file name is safe
        const safeName = path.basename(fileName);
        await fs.writeFile(safeName, "");
        return `File ${safeName} created successfully`;
      } catch (error) {
        return `Error creating file: ${error.message}`;
      }
    },
  });

  // Tool to write content to a file
  const writeFileTool = new DynamicTool({
    name: "write_to_file",
    description:
      "Writes content to a file. Input should be in format 'filename|||content'",
    func: async (input: string) => {
      try {
        const [fileName, content] = input.split("|||");
        const safeName = path.basename(fileName);
        await fs.writeFile(safeName, content);
        return `Content written to ${safeName} successfully`;
      } catch (error) {
        return `Error writing to file: ${error.message}`;
      }
    },
  });

  // Tool to read file content
  const readFileTool = new DynamicTool({
    name: "read_file",
    description: "Reads content from a file",
    func: async (fileName: string) => {
      try {
        const safeName = path.basename(fileName);
        const content = await fs.readFile(safeName, "utf-8");
        return content;
      } catch (error) {
        return `Error reading file: ${error.message}`;
      }
    },
  });

  const tools = [createFileTool, writeFileTool, readFileTool];

  const executor = await initializeAgentExecutorWithOptions(tools, model, {
    agentType: "chat-conversational-react-description",
    verbose: true,
  });

  return executor;
}

// Example usage
async function main() {
  const agent = await createFileOperationsAgent();

  // Create a new file
  const createResult = await agent.call({
    input: "Create a new file called example.txt",
  });
  console.log(createResult.output);

  // Write content to the file
  const writeResult = await agent.call({
    input: "Write 'Hello, World!' to example.txt",
  });
  console.log(writeResult.output);

  // Read the file content
  const readResult = await agent.call({
    input: "What's the content of example.txt?",
  });
  console.log(readResult.output);
}

// Example with more complex content generation
async function createDocumentationFile() {
  const agent = await createFileOperationsAgent();

  // Create and write documentation
  const result = await agent.call({
    input:
      "Create a file called api-docs.md and write a brief API documentation for a user registration endpoint. Include sections for request format, response format, and error handling.",
  });
  console.log(result.output);
}

// Run examples
main().catch(console.error);
createDocumentationFile().catch(console.error);
