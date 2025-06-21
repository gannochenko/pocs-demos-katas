import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { join } from 'path';
import { ApolloServerPluginLandingPageLocalDefault } from '@apollo/server/plugin/landingPage/default';
import * as depthLimit from 'graphql-depth-limit';
import { AuthorsModule } from '../authors/authors.module';
import { PostsModule } from '../posts/posts.module';
import { ApolloServerPluginComplexity } from 'src/util/graphql.complexityPlugin';
import { HealthModule } from 'src/health/health.module';

@Module({
  imports: [
    AuthorsModule,
    PostsModule,
    HealthModule,
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      typePaths: ['./**/*.graphql'],
      definitions: {
        path: join(process.cwd(), 'src/graphql.ts'),
        outputAs: 'class',
      },
      playground: false,
      plugins: [
        ApolloServerPluginLandingPageLocalDefault(), // todo: enable only in development
        ApolloServerPluginComplexity(),
      ],
      validationRules: [depthLimit(10)],
      formatError: (error) => {
        // don't disclose too much info to the client
        return {
          message: error.message,
          locations: error.locations,
          path: error.path,
        };
      },
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
