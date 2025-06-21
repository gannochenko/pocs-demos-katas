import {
  BaseContext,
  GraphQLRequestContext,
  GraphQLRequestListener,
} from '@apollo/server';
import { GraphQLError } from 'graphql';
import {
  fieldExtensionsEstimator,
  getComplexity,
  simpleEstimator,
} from 'graphql-query-complexity';

export function ApolloServerPluginComplexity<TContext extends BaseContext>() {
  return {
    // eslint-disable-next-line @typescript-eslint/require-await
    async requestDidStart(): Promise<GraphQLRequestListener<any>> {
      const maxComplexity = 1000; // todo: get from config

      return {
        // eslint-disable-next-line @typescript-eslint/require-await
        didResolveOperation: async (
          requestContext: GraphQLRequestContext<TContext>,
        ) => {
          const { request, document, schema } = requestContext || {};

          const complexity = getComplexity({
            schema,
            operationName: request.operationName,
            query: document!, // todo: fix this
            variables: request.variables,
            estimators: [
              // Using fieldExtensionsEstimator is crucial to read complexity values from `@Field({ complexity: X })`
              fieldExtensionsEstimator(),
              // simpleEstimator assigns a default complexity to fields that don't have a specific complexity defined
              simpleEstimator({ defaultComplexity: 1 }),
            ],
          });

          if (complexity > maxComplexity) {
            throw new GraphQLError(
              `Query is too complex: ${complexity}. Maximum allowed complexity: ${maxComplexity}`,
              {
                extensions: {
                  code: 'GRAPHQL_VALIDATION_FAILED',
                  http: { status: 400 },
                },
              },
            );
          }
        },
      };
    },
  };
}

// @Plugin()
// export class ComplexityPlugin implements ApolloServerPlugin<TContext extends BaseContext> {
//   constructor(private gqlSchemaHost: GraphQLSchemaHost) {}

//   // eslint-disable-next-line @typescript-eslint/require-await
//   async requestDidStart(): Promise<GraphQLRequestListener<any>> {
//     const { schema } = this.gqlSchemaHost;
//     const maxComplexity = 20; // Define your maximum allowed complexity here

//     return {
//       // eslint-disable-next-line @typescript-eslint/require-await
//       async didResolveOperation(requestContext: GraphQLRequestContext<TContext>) {
//         const { request, document } = requestContext || {};

//         const complexity = getComplexity({
//           schema,
//           operationName: request.operationName,
//           query: document!, // todo: fix this
//           variables: request.variables,
//           estimators: [
//             // Using fieldExtensionsEstimator is crucial to read complexity values from `@Field({ complexity: X })`
//             fieldExtensionsEstimator(),
//             // simpleEstimator assigns a default complexity to fields that don't have a specific complexity defined
//             simpleEstimator({ defaultComplexity: 1 }),
//           ],
//         });

//         if (complexity > maxComplexity) {
//           throw new GraphQLError(
//             `Query is too complex: ${complexity}. Maximum allowed complexity: ${maxComplexity}`,
//             {
//               extensions: {
//                 code: 'GRAPHQL_VALIDATION_FAILED',
//                 http: { status: 400 },
//               },
//             },
//           );
//         }
//         console.log(`Query Complexity: ${complexity}`);
//       },
//     };
//   }
// }
