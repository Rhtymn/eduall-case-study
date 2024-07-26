import { GetAllProductQuery } from "../types/requests/product";
import { APIResponse } from "../types/responses";
import { ProductResponseWithMeta } from "../types/responses/product";
import apiService from "../utils/apiService";
import { wrapError } from "../utils/exception";

export const getAllProducts = async (
  q?: GetAllProductQuery
): Promise<APIResponse<ProductResponseWithMeta> | undefined> => {
  try {
    const queryParams = new URLSearchParams();
    if (q?.limit) {
      queryParams.append("limit", `${q.limit}`);
    }

    if (q?.page) {
      queryParams.append("page", `${q.page}`);
    }

    if (q?.search_term) {
      queryParams.append("search_term", q.search_term);
    }

    const res = await apiService.get<
      APIResponse<ProductResponseWithMeta> | undefined
    >(`products?${queryParams.toString()}`);

    return res;
  } catch (e) {
    throw await wrapError(e);
  }
};
