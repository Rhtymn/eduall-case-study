import { PageInfo } from ".";

export interface Product {
  id: number;
  brand: string;
  model: string;
  screen_size: string;
  color: string;
  harddisk: string;
  cpu: string;
  ram: string;
  os: string;
  special_features: string;
  graphics: string;
  graphics_coprocessor: string;
  cpu_speed: string;
  rating: number;
  price: string;
}

export interface ProductResponseWithMeta {
  products: Product[];
  meta: PageInfo;
}
