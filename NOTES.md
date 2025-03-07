# NOTES
## The schema issue
TODO: Describe relational vs document


## Dates
All dates are formatted according to ISO 8601 

## Rationale
I started by taking notes on the structure of the documents. 

These documents have too much variability to create a strong schema. That is to say, there is no configuration of columns that is satisfactory for each and every document. This makes a relational data model a poor choice for encoding. Instead, I opted to use a document database. 

*"If you're not exactly sure how your data is structured at this point, a document database is probably the best place to start"* - Jeff Delaney, Google Developer Expert 

Document databases provide the flexibility demanded by the project. It is also easy to convert a document representation to a relational one. 

I would look at the longest documents first. I figure there is a higher probability I would find edge cases. 

## Groups
### Owyhee
1.
2.
3. Toy Mountain
4. South Mountain
5. Morgan
