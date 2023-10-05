const fs = require('fs');
const yaml = require('yaml');

const args = process.argv.splice(2);

function loadYaml(p) {
  try {
    return yaml.parse(fs.readFileSync(p, 'utf8'));
  } catch (err) {
    console.error(err);
    process.exit(1);
  }
}

function visit(o, fn) {
  let toVisit = [{path: null, value: o}];
  while (toVisit.length > 0) {
    const {path, value} = toVisit.pop()
    if (value == null) {
      continue;
    } 
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

let actual = loadYaml(args[0]);
for (let i = 1; i < args.length; i++) {
  let toMerge = loadYaml(args[i]);
  visit(toMerge, function (k, v) {
    let path = actual;
    k.forEach((key, i) => {
      path[key] ||= {};
      if (i == (k.length - 1)) {
        path[key] = v;
      }
      path = path[key];
    });
  });
}

process.stdout.write(yaml.stringify(actual));
