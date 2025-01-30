from fastapi import FastAPI
from app.routes.hello import router as hello_router

app = FastAPI()

# Include the Hello World router
app.include_router(hello_router)

# Run the app using `uvicorn` or any ASGI server.
