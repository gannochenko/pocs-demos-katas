import { Module } from '@nestjs/common';
import { AuthorsResolver } from './authors.resolver';
import { AuthorsService } from './authors.service';
import { PostsService } from 'src/posts.service';

@Module({
  imports: [],
  providers: [AuthorsService, PostsService, AuthorsResolver],
})
export class AuthorsModule {}
