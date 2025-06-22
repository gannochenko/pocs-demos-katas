import { Module } from '@nestjs/common';
import { PostsResolver } from './posts.resolver';
import { PostsService } from './posts.service';
import { AuthorsService } from '../authors/authors.service';
import { HttpModule } from '@nestjs/axios';

@Module({
  imports: [HttpModule.register({})],
  providers: [PostsService, AuthorsService, PostsResolver],
})
export class PostsModule {}
