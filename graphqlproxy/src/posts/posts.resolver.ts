import { Parent, Resolver, ResolveField } from '@nestjs/graphql';
import { AuthorsService } from '../authors/authors.service';

@Resolver('Post')
export class PostsResolver {
  constructor(private authorsService: AuthorsService) {}

  @ResolveField()
  author(@Parent() post: { authorId: number }) {
    const { authorId } = post;
    return this.authorsService.findOneById(authorId);
  }
}
