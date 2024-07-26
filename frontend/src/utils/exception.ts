import { HttpStatus } from "../constants/api";
import {
  BadRequest,
  InternalServerError,
  ResponseError,
} from "../exceptions/apiExceptions";

export const wrapError = async (e: unknown): Promise<Error> => {
  if (e instanceof ResponseError) {
    const res = await e.response.json();
    switch (e.response.status) {
      case HttpStatus.BadRequest:
        return new BadRequest(res.message);
      default:
        return new InternalServerError(res.message);
    }
  }
  return new InternalServerError("something went wrong");
};
