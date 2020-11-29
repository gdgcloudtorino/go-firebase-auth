# go-firebase-auth


#### Deploy 

gcloud functions deploy HelloHTTP --region=europe-west3 --memory=128MB --runtime go113 --trigger-http --allow-unauthenticated

gcloud functions deploy MockLogin --region=europe-west3 --memory=128MB --runtime go113 --trigger-http


"gcloud alpha functions add-iam-policy-binding MockLogin --region=europe-west3 --member=allUsers --role=roles/cloudfunctions.invoker"
