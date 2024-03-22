# Intermediate docker image to build the bundle in and install dependencies
FROM node:21-alpine as build

# Set the working directory to /usr/src/app
WORKDIR /usr/src/app

# Copy the package.json and package-lock.json over in the intermedate "build" image
COPY ./package.json ./

# Install the dependencies
RUN npm i

# Copy the source code into the build image
COPY ./ ./

# Build the project
RUN npm run build

# Pull the same Node image and use it as the final (production image)
FROM node:21-alpine as production

# Set the working directory to /usr/src/app
WORKDIR /usr/src/app

# Only copy the results from the build over to the final image
# We do this to keep the final image as small as possible
COPY --from=build /usr/src/app/ .


EXPOSE 5173

# Start the application
CMD [ "npm", "run", "start"]
