import { ResponseError } from "../exceptions/apiExceptions";

class ApiService {
  constructor(private baseUrl: string) {}

  private async request<T>(
    resource: string,
    method: string,
    req?: Omit<RequestInit, "method">
  ): Promise<T | undefined> {
    const res = await fetch(`${this.baseUrl}/${resource}`, {
      ...req,
      method: method,
      headers: {
        "Content-Type": "application/json",
        ...req?.headers,
      },
    });

    if (res.status >= 400) {
      throw new ResponseError("something went wrong", res);
    }

    return await res.json();
  }

  async get<T>(
    resource: string,
    req?: Omit<RequestInit, "method">
  ): Promise<T | undefined> {
    return await this.request(resource, "GET", req);
  }
}

const apiService = new ApiService("http://localhost:8080/api/v1");
export default apiService;
