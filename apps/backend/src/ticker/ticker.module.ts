import { Module } from '@nestjs/common';
import { TickerService } from './ticker.service';
import { TickerController } from './ticker.controller';
import { ClientsModule } from '@nestjs/microservices';
import { tickerClientGrpc } from 'src/grpc/ticker.client';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'TICKER_PACKAGE',
        ...tickerClientGrpc,
      },
    ]),
  ],
  providers: [TickerService],
  controllers: [TickerController],
})
export class TickerModule {}
