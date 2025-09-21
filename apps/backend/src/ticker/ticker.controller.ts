import { Controller, Get, Query, Sse } from '@nestjs/common';
import { TickerService } from './ticker.service';
import { Observable } from 'rxjs';
import { TickerUpdate } from './ticker.interface';

@Controller('ticker')
export class TickerController {
  constructor(private readonly tickerService: TickerService) {}

  @Get('stream')
  @Sse()
  stream(@Query('symbol') symbol: string): Observable<TickerUpdate> {
    console.log(symbol);
    return this.tickerService.streamTicker(symbol.toLowerCase());
  }
}
