# Fetcher

Fetch news from a entrance url, filter update time in the range, and storage them into a json. So, Microservice can read this fetched data directly.

## Skeleton

- fetcher.go: Fetch action implemented here
- links.go: Fetch links from an entrance url.
- article.go: Domain Object here to implement functions that microservice needed and invoked, such as Get, List and Search.
- db.go: This is for fetched data persistence, load and save data directly.

## Fetch kernal methods

There are kernal methods used to modified to match and fetch better, you can clone and rewrite the parts to implement your microservice.  
  
1. `func (a *Article) fetchTitle() (string, error) {}`  
2. `func (a *Article) fetchUpdateTime() (*timestamppb.Timestamp, error) {}`  
3. `func (a *Article) fetchContent() (string, error) {}`
