export interface Parameter {
  name: string;
  type: string;
  required: boolean;
  location: string;
  description: string;
  example?: string;
}

export interface ResponseField {
  name: string;
  type: string;
  description: string;
}

export interface ApiError {
  status: number;
  code?: string;
  description: string;
}

export interface Endpoint {
  name: string;
  method: string;
  path: string;
  description: string;
  parameters?: Parameter[];
  request_example?: string;
  response_example?: string;
  response_fields?: ResponseField[];
  errors?: ApiError[];
}

export interface ApiDoc {
  api_id: string;
  api_name: string;
  description: string;
  base_url: string;
  endpoints: Endpoint[];
}
