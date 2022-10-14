declare class OSPResponse {
  version: string;
  status: number;
  reason: string;
  headers: Map<string, string>;
  body?: string;
}
