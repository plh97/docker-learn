FROM node
ADD . /code
WORKDIR /code
RUN npm install
RUN yarn global add nodemon
CMD ["yarn", "start"]