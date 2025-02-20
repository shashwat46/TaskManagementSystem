interface TaskFiltersProps {
    statusFilter: string;
    priorityFilter: string;
    sortBy: string;
    onStatusChange: (status: string) => void;
    onPriorityChange: (priority: string) => void;
    onSortChange: (sort: string) => void;
  }
  
  export default function TaskFilters({
    statusFilter,
    priorityFilter,
    sortBy,
    onStatusChange,
    onPriorityChange,
    onSortChange,
  }: TaskFiltersProps) {
    return (
      <div className="flex flex-wrap gap-4 mb-6">
        <div className="flex-1 min-w-[200px]">
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Status
          </label>
          <select
            value={statusFilter}
            onChange={(e) => onStatusChange(e.target.value)}
            className="w-full rounded-md border border-gray-300 p-2"
          >
            <option value="">All Status</option>
            <option value="pending">Pending</option>
            <option value="in_progress">In Progress</option>
            <option value="completed">Completed</option>
          </select>
        </div>
  
        <div className="flex-1 min-w-[200px]">
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Priority
          </label>
          <select
            value={priorityFilter}
            onChange={(e) => onPriorityChange(e.target.value)}
            className="w-full rounded-md border border-gray-300 p-2"
          >
            <option value="">All Priorities</option>
            <option value="low">Low</option>
            <option value="medium">Medium</option>
            <option value="high">High</option>
          </select>
        </div>
  
        <div className="flex-1 min-w-[200px]">
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Sort By
          </label>
          <select
            value={sortBy}
            onChange={(e) => onSortChange(e.target.value)}
            className="w-full rounded-md border border-gray-300 p-2"
          >
            <option value="due_date_asc">Due Date (Earliest)</option>
            <option value="due_date_desc">Due Date (Latest)</option>
            <option value="priority_high">Priority (High-Low)</option>
            <option value="priority_low">Priority (Low-High)</option>
          </select>
        </div>
      </div>
    );
  }
  