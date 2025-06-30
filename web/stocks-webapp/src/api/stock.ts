import axiosInstance from './index'

export async function getStocksFromApi(
  offset: number,
  sortBy: string = 'time',
  sortDir: 'asc' | 'desc' = 'desc',
  searchTerm?: string,
  searchBy: string = 'company', // search by company
) {
  try {
    const response = await axiosInstance.get('/stocks', {
      params: { limit: 20, offset, sort_by: sortBy, sort_dir: sortDir },
    })
    return response.data
  } catch (error) {
    console.error('Error fetching stocks:', error)
    throw error
  }
}

