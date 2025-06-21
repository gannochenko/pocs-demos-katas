export class PostsService {
  findAll(filters: { authorId: number }) {
    // TODO: Implement actual logic
    return [
      {
        id: 1,
        title: `Post for author ${filters.authorId}`,
        authorId: filters.authorId,
      },
    ];
  }
}
