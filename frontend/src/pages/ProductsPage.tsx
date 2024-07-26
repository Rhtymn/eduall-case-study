import React, { ChangeEvent } from "react";
import { Product } from "../types/responses/product";
import { getAllProducts } from "../services/product";
import { BadRequest, InternalServerError } from "../exceptions/apiExceptions";
import { PageInfo } from "../types/responses";
import ProductTable from "../components/tables/ProductTable";
import Paginator from "../components/ui/Paginator";
import { useSearchParams } from "react-router-dom";
import { debounce } from "../utils/helper";

const ProductsPage = () => {
  const [products, setProducts] = React.useState<Product[]>([]);
  const [meta, setMeta] = React.useState<PageInfo | undefined>(undefined);
  const [searchParams, setSearchParams] = useSearchParams();
  const [error, setError] = React.useState<string>("");

  const pageParam = searchParams.get("page");
  let page: number = 1;
  if (!pageParam) {
    searchParams.set("page", "1");
  }

  if (pageParam) {
    if (!isNaN(+pageParam)) {
      page = +pageParam;
      searchParams.set("page", pageParam);
    } else {
      searchParams.set("page", "1");
    }
  }

  const searchTermParam = searchParams.get("search_term");
  let searchTerm: string | undefined = undefined;
  if (searchTermParam) {
    searchTerm = searchTermParam;
  }

  React.useEffect(() => {
    const getProducts = async () => {
      try {
        const res = await getAllProducts({
          page: page,
          limit: 10,
          search_term: searchTerm,
        });
        setProducts(res?.data?.products ?? []);
        setMeta(res?.data?.meta ?? undefined);
      } catch (e) {
        if (e instanceof BadRequest) {
          setError(e.message);
        } else if (e instanceof InternalServerError) {
          setError(e.message);
        } else {
          setError("something went wrong");
        }
      }
    };

    getProducts();
  }, [page, searchTerm]);

  const onPrev = () => {
    if (page === 1) return;
    searchParams.set("page", `${page - 1}`);
    setSearchParams(searchParams);
  };

  const onNext = () => {
    if (meta?.current_page === meta?.page_count || meta?.page_count === 0)
      return;
    searchParams.set("page", `${page + 1}`);
    setSearchParams(searchParams);
  };

  const debouncedOnChange = debounce((value: string) => {
    searchParams.set("search_term", value);
    setSearchParams(searchParams);
  }, 300);

  const onChange = (e: ChangeEvent<HTMLInputElement>) => {
    debouncedOnChange(e.currentTarget.value);
  };

  return (
    <main className="px-2 py-4">
      <div className="flex flex-col gap-y-4">
        <h1 className="text-xl font-bold">Product List</h1>
        <SearchInput onChange={onChange} />
        <ProductTable products={products} meta={meta} error={error} />
        {!error && <Paginator meta={meta} onPrev={onPrev} onNext={onNext} />}
      </div>
    </main>
  );
};

const SearchInput = ({ onChange }: React.ComponentProps<"input">) => {
  return (
    <input
      onChange={onChange}
      className="border-2 rounded-md p-1 px-2"
      placeholder="Search brand or model..."
    />
  );
};

export default ProductsPage;
