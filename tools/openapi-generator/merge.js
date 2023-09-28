const fs = require('fs');
const yaml = require('js-yaml');

var atlasApi, atlasApiChanges;

try {
  atlasApi = yaml.load(fs.readFileSync('config/atlas-api.yaml', 'utf8'));
} catch (err) {
    console.error(err);
    process.exit(1);
}

try {
  atlasApiChanges = yaml.load(fs.readFileSync('config/atlas-api-changes.yaml', 'utf8'));
} catch (err) {
    console.error(err);
    process.exit(1);
}

function visit(o, fn) {
  let toVisit = [{path: null, value: o}];
  while (toVisit.length > 0) {
    const {path, value} = toVisit.pop()
    if (typeof value !== 'object') {
      fn(path, value);
      continue;
    }

    Object.entries(value).forEach(entry => {
      const [key, value] = entry;
      let newPath = [key];
      if (path != null) {
        newPath = [...path, key];
      } 
      toVisit.push({path: newPath, value: value});
    });
  }
}

visit(atlasApiChanges, function (k, v) {
  let path = atlasApi;
  k.forEach((key, i) => {
    path[key] ||= {};
    if (i == (k.length - 1)) {
      path[key] = v;
    }
    path = path[key];
  });
});

try {
  fs.writeFileSync('config/atlas-api-transformed.yaml', yaml.dump(atlasApi), 'utf8')
} catch (err) {
    console.error(err);
    process.exit(1);
}