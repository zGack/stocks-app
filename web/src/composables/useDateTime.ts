export function useDateTime() {
  function formatDate(dateStr: string) {
    const d = new Date(dateStr)
    return d.toLocaleString()
  }

  return {
    formatDate,
  }
}
