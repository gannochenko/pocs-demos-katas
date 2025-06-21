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
import { complexityRule } from 'src/util/graphql.complexityRule';

@Module({
  imports: [
    AuthorsModule,
    PostsModule,
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      typePaths: ['./**/*.graphql'],
      definitions: {
        path: join(process.cwd(), 'src/graphql.ts'),
        outputAs: 'class',
      },
      playground: false,
      plugins: [ApolloServerPluginLandingPageLocalDefault()], // todo: enable only in development
      validationRules: [depthLimit(10)],
    }),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
