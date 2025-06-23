import { Author } from './dto/author';
import { HttpService } from '@nestjs/axios';
import { JsonPlaceholderUser } from 'src/jsonplaceholder/jsonplaceholder.types';
import { retryRequest } from 'src/util/retryRequest';
import { Injectable } from '@nestjs/common';

@Injectable()
export class AuthorsService {
  constructor(private readonly httpService: HttpService) {}

  async findOneById(id: number): Promise<Author> {
    const data = await retryRequest(
      this.httpService.get<JsonPlaceholderUser>(
        `https://jsonplaceholder.typicode.com/users/${id}`,
        {
          maxRedirects: 5,
        },
      ),
    );

    return {
      id,
      firstName: data.name,
      lastName: data.username,
    };
  }
}
