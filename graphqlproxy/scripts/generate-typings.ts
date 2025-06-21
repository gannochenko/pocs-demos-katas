import { GraphQLDefinitionsFactory } from '@nestjs/graphql';
import { join } from 'path';

const definitionsFactory = new GraphQLDefinitionsFactory();
definitionsFactory
  .generate({
    typePaths: [join(process.cwd(), 'schema.graphql')],
    path: join(process.cwd(), 'src/graphql.ts'),
    outputAs: 'class',
  })
  .catch((err) => {
    console.error(err);
  });
