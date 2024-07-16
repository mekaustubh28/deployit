FROM node:18.19.0

WORKDIR /app

COPY . ./

RUN npm install

ENV VITE_BASE_URL=http://localhost:8001

EXPOSE 5173

CMD ["npm", "run", "dev"]
