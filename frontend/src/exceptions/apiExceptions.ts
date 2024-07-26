export class ResponseError extends Error {
  response: Response;

  constructor(message: string, res: Response) {
    super(message);
    this.response = res;
  }
}

export class InternalServerError extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ErrInternalServerError";
  }
}

export class BadRequest extends Error {
  constructor(message: string) {
    super(message);
    this.name = "ErrBadRequest";
  }
}
