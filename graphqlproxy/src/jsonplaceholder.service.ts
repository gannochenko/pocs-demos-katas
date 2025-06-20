import { Injectable } from '@nestjs/common';

@Injectable()
export class JsonPlaceholderService {
  getHello(): string {
    return 'Hello World!';
  }
}
