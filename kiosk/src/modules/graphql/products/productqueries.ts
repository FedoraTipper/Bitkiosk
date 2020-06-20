export default {
  loadActiveProducts: `
    query {
          loadActiveProducts{
            SKU
            name
            description
            shortDescription
            price
            stock
            reviewCount
            rating
            startDate
            endDate
            createdAt
          }
    } 
    `
}