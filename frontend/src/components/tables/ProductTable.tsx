import { PageInfo } from "../../types/responses";
import { Product } from "../../types/responses/product";

const ProductTable = ({ products, meta, error }: ProductTableProps) => {
  return (
    <>
      <div className="overflow-x-scroll border-gray-400">
        <table className="table-auto overflow-x-scroll w-full border">
          <ProductTableHeader />
          {products.length > 0 && (
            <tbody>
              {products.map((p, idx) => (
                <ProductRow key={idx} idx={idx} product={p} meta={meta} />
              ))}
            </tbody>
          )}
        </table>
      </div>
      {products.length === 0 && !error && (
        <p className="text-center">Product not found</p>
      )}
      {error && <p className="text-center">{error}</p>}
    </>
  );
};

const ProductTableHeader = () => {
  return (
    <thead>
      <tr className="bg-gray-400">
        <th className="text-center min-w-[40px] py-2">No</th>
        <th className="min-w-[100px] text-start">Brand</th>
        <th className="min-w-[160px] text-start">Model</th>
        <th className="min-w-[110px] text-start">Screen Size</th>
        <th className="min-w-[100px] text-start">Color</th>
        <th className="min-w-[100px] text-start">Harddisk</th>
        <th className="min-w-[150px] text-start">CPU</th>
        <th className="min-w-[80px] text-start">RAM</th>
        <th className="min-w-[150px] text-start">OS</th>
        <th className="min-w-[140px] text-start">Special Features</th>
        <th className="min-w-[90px] text-start">Graphics</th>
        <th className="min-w-[180px] text-start">Graphic Coprocessor</th>
        <th className="min-w-[100px] text-start">Cpu Speed</th>
        <th className="min-w-[80px] text-start">Rating</th>
        <th className="min-w-[100px] text-start">Price</th>
      </tr>
    </thead>
  );
};

const ProductRow = ({
  idx,
  product,
  meta,
}: {
  idx: number;
  meta?: PageInfo;
  product: Product;
}) => {
  if (!meta) return;

  return (
    <tr className="even:bg-gray-200">
      <td className="text-center py-2">
        {(meta.current_page - 1) * meta.items_per_page + (idx + 1)}
      </td>
      <td>{product.brand}</td>
      <td>{product.model}</td>
      <td>{product.screen_size}</td>
      <td>{product.color}</td>
      <td>{product.harddisk}</td>
      <td>{product.cpu}</td>
      <td>{product.ram}</td>
      <td>{product.os}</td>
      <td>{product.special_features}</td>
      <td>{product.graphics}</td>
      <td>
        {product.graphics_coprocessor == ""
          ? "-"
          : product.graphics_coprocessor}
      </td>
      <td>{product.cpu_speed}</td>
      <td>{product.rating}</td>
      <td>{product.price}</td>
    </tr>
  );
};

interface ProductTableProps {
  products: Product[];
  meta?: PageInfo;
  error?: string;
}

export default ProductTable;
