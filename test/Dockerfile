FROM ruby:2.7.4

WORKDIR /code
COPY . /code
RUN bundle install

EXPOSE 3000
EXPOSE 9292
EXPOSE 9393

CMD ["bundle", "exec", "puma"]
