import React from "react";

type OptionConfig = {
  label: string;
  colorClass: string;
  hoverClass: string;
};

const OPTIONS: OptionConfig[] = [
  {
    label: "反対",
    colorClass: "bg-red-400",
    hoverClass: "hover:bg-white hover:border-red-600",
  },
  {
    label: "やや反対",
    colorClass: "bg-red-300",
    hoverClass: "hover:bg-white hover:border-red-300",
  },
  {
    label: "中立",
    colorClass: "bg-gray-300",
    hoverClass: "hover:bg-white hover:border-gray-400",
  },
  {
    label: "やや賛成",
    colorClass: "bg-blue-200",
    hoverClass: "hover:bg-white hover:border-blue-200",
  },
  {
    label: "賛成",
    colorClass: "bg-blue-400",
    hoverClass: "hover:bg-white hover:border-blue-500",
  },
];

const BASE_BUTTON_CLASSES =
  "flex-1 px-4 py-6 border-2 border-transparent rounded-xl cursor-pointer text-lg font-bold text-center whitespace-nowrap text-black transition-all duration-200";

interface SelectApproveOrRejectProps {
  onSelect: (choice: string) => void;
  className?: string;
}

export default function SelectApproveOrReject({
  onSelect,
  className = "",
}: SelectApproveOrRejectProps) {
  return (
    <div
      className={`flex gap-2 justify-between w-full max-w-5xl mx-auto ${className}`}
    >
      {OPTIONS.map((option) => (
        <button
          key={option.label}
          className={`${BASE_BUTTON_CLASSES} ${option.colorClass} ${option.hoverClass}`}
          onClick={() => onSelect(option.label)}
        >
          {option.label}
        </button>
      ))}
    </div>
  );
}

// Export constants for reuse if needed
export { OPTIONS, type OptionConfig };
