FROM retypeapp/retype:3-alpine AS build

# Install ICU libraries to fix globalization-invariant mode error
RUN apk add --no-cache icu-libs

# Set environment variable to disable invariant globalization mode
ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=false

WORKDIR /site

COPY . .

RUN retype build /site


FROM nginx:alpine AS runtime

LABEL org.opencontainers.image.title="parsley-docs" \
      org.opencontainers.image.description="Parsley documentation site built with Retype and served by NGINX" \
      org.opencontainers.image.source="https://github.com/matzefriedrich/parsley-docs"

COPY --from=build /site/.retype /usr/share/nginx/html

EXPOSE 80
