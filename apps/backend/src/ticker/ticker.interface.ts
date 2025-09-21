import { Observable } from 'rxjs';

export interface TickerRequest {
  symbol: string;
}

export interface TickerUpdate {
  symbol: string;

  volume: string;

  high: string;
}

export interface TickerService {
  StreamTicker(request: TickerRequest): Observable<TickerUpdate>;
}
