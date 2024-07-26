import React from "react";
import { PageInfo } from "../../types/responses";

const Paginator = ({ meta, onNext, onPrev }: PaginatorProps) => {
  if (meta?.page_count === 0) return;
  return (
    <div className="flex gap-x-4 items-center w-full justify-center">
      <ActionButton
        onClick={onPrev}
        disabled={meta?.current_page === 1}
        className={`${
          meta?.current_page === 1 ? "opacity-30 cursor-not-allowed" : ""
        }`}
      >
        Prev
      </ActionButton>
      <p>
        {meta?.current_page} of {meta?.page_count}
      </p>
      <ActionButton
        onClick={onNext}
        disabled={meta?.current_page === meta?.page_count}
        className={`${
          meta?.current_page === meta?.page_count
            ? "opacity-30 cursor-not-allowed"
            : ""
        }`}
      >
        Next
      </ActionButton>
    </div>
  );
};

interface PaginatorProps {
  meta?: PageInfo;
  onNext: () => void;
  onPrev: () => void;
}

const ActionButton = ({
  className,
  onClick,
  children,
}: React.ComponentProps<"button">) => {
  return (
    <button
      className={`bg-gray-500 text-white px-3 py-1 rounded-md ${className}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Paginator;
