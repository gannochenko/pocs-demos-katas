import { Author } from './dto/author';

export class AuthorsService {
  findOneById(id: number): Author {
    return { id, firstName: 'John', lastName: 'Doe' };
  }
}
