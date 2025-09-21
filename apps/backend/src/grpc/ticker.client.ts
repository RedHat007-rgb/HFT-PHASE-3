import { ClientOptions } from '@grpc/grpc-js';
import { Transport } from '@nestjs/microservices';
import { join } from 'path';

export const tickerClientGrpc: ClientOptions = {
  transport: Transport.GRPC,
  options: {
    package: 'ticker',
    protoPath: join(__dirname, '../../../../packages/proto/ticker.proto'),
    url: 'localhost:50052',
  },
};
