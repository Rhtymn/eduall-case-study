export interface APIResponse<T> {
  message: string;
  data?: T;
}

export interface PageInfo {
  current_page: number;
  items_per_page: number;
  item_count: number;
  page_count: number;
}
