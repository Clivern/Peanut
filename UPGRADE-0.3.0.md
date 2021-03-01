UPGRADE FROM 0.1.0 and 0.2.0 to 0.3.0
=====================================

Now Peanut supports deploying certain versions. it is not required when you call the API since it will use the default (latest version). You will need to do the following in order to update

* Update config file `containerization` property.

```yaml
    # Containerization runtime (supported docker)
    containerization:
        driver: ${PEANUT_CONTAINERIZATION_DRIVER:-docker}

        # Clean up stale images, volumes and networks
        autoClean: ${PEANUT_CONTAINERIZATION_AUTO_CLEAN:-true}

        # Time to cache docker images tags
        cacheTagsTimeInMinutes: ${PEANUT_CONTAINERIZATION_CACHE_TIME:-10080}
```
