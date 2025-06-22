import { Injectable } from '@nestjs/common';

@Injectable()
export class PostsService {
  findAll(filters: { authorId: number }) {
    // TODO: Implement actual logic

    // throw new Error('Could not fetch posts');

    return [
      {
        id: 1,
        title: `Post for author ${filters.authorId}`,
        authorId: filters.authorId,
      },
    ];
  }
}
