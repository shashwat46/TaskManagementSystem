"use client";

export default function DashboardPage() {
  return (
    <div className="bg-white shadow rounded-lg p-6">
      <h2 className="text-2xl font-bold mb-4">Welcome to Your Dashboard</h2>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-blue-50 p-4 rounded-lg">
          <h3 className="font-semibold text-lg mb-2">Pending Tasks</h3>
          <p className="text-3xl font-bold text-blue-600">5</p>
        </div>
        <div className="bg-green-50 p-4 rounded-lg">
          <h3 className="font-semibold text-lg mb-2">Completed Tasks</h3>
          <p className="text-3xl font-bold text-green-600">12</p>
        </div>
        <div className="bg-purple-50 p-4 rounded-lg">
          <h3 className="font-semibold text-lg mb-2">Total Tasks</h3>
          <p className="text-3xl font-bold text-purple-600">17</p>
        </div>
      </div>
    </div>
  );
}
