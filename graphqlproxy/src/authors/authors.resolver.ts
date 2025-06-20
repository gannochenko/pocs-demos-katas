import { Parent, Resolver, Query, Args, ResolveField } from '@nestjs/graphql';
import { AuthorsService } from './authors.service';
import { PostsService } from '../posts.service';

@Resolver('Author')
export class AuthorsResolver {
  constructor(
    private authorsService: AuthorsService,
    private postsService: PostsService,
  ) {}

  @Query()
  author(@Args('id') id: number) {
    return this.authorsService.findOneById(id);
  }

  @ResolveField()
  posts(@Parent() author: { id: number }) {
    const { id } = author;
    return this.postsService.findAll({ authorId: id });
  }
}
