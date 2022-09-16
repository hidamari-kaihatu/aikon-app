FROM node:16.17.0
WORKDIR /usr/src/app
# COPY package*.json ./
# RUN npm install -g npm@latest && npm install create-next-app

COPY . .
RUN npm install

EXPOSE 3000

CMD ["npm", "run", "dev"] 
