export function Spinner() {
  return (
    <div className="fixed inset-0 flex items-center justify-center bg-white bg-opacity-80">
      <div className="w-12 h-12 border-4 border-gray-200 rounded-full border-t-green-500 animate-spin" />
    </div>
  );
}
