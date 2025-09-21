import { Inject, Injectable, OnModuleInit } from '@nestjs/common';
import {
  TickerService as TickerGrpcService,
  TickerUpdate,
} from './ticker.interface';
import type { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';

@Injectable()
export class TickerService implements OnModuleInit {
  private tickerGrpc: TickerGrpcService;

  constructor(@Inject('TICKER_PACKAGE') private client: ClientGrpc) {}
  onModuleInit() {
    this.tickerGrpc =
      this.client.getService<TickerGrpcService>('TickerService');
  }

  streamTicker(symbol: string): Observable<TickerUpdate> {
    console.log('in service');
    return this.tickerGrpc.StreamTicker({ symbol: symbol });
  }
}
