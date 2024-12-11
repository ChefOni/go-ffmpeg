# PROJECT DESCRIPTION
It is a transcoder of some sorts that accepts video , reduces the size and uploads it to a storage bucket of the user's choice 

## LIMITS
It might be limited to only some services like AWS S3 bucket , Cloudinary and Cloudflare's Storage Bucket
It does not take a video of more than a certain size 

## ROUTES
"/" - Welcomes the user 
"/health" - Check the health of the server
"/upload" - Upload the video resource
"/check-status" - Returns the status of the video resource

