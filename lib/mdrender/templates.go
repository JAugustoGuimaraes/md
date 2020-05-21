package mdrender

const baseIndexHtml = `<!DOCTYPE html>
<head>
  <title>{{.Title}}</title>
  <meta charset="UTF-8">
</head>
<body>
  {{.Content}}
  
  <!-- Websocket script -->
  %s

  <!-- CSS switcher script -->
  %s
</body>`

const websocketJs = `<script>
  const socket = new WebSocket("ws://localhost:8080/ws");

  socket.onopen = function() {
    console.log("Websocket connected");
  };

  socket.onmessage = function (e) {
    window.location.reload();
  };
</script>`

// Inspiration code taken from https://github.com/dohliam/dropin-minimal-css
const cssSwitchJs = `<script>
  const frameworks = "a11yana,awsm,bahunya,bare,base,bullframe,bulma,caiuss,caramel,cardinal,chota,clmaterial,codify,comet,concise,concrete,flat-ui,fluidity,furtive,generic,github-markdown,hack,holiday,html-starterkit,hyp,kathamo,koochak,kraken,kube,latex,lemon,lit,lotus,markdown,marx,materialize,mercury,milligram,min,mini,mobi,motherplate,mu,mui,mvp,new,no-class,normalize,oh-my-css,paper,papier,pavilion,picnic,preface,primer,pure,sakura,sanitize,semantic-ui,shoelace,siimple,simple,skeleton,skeleton-framework,skeleton-plus,snack,spectre,style,stylize,tachyons,tacit,tent,thao,vanilla,vital,water,wing,writ,yamb,yorha,ads-gazette,ads-medium,ads-notebook,ads-tufte,boot-cerulean,boot-cosmo,boot-cyborg,boot-darkly,boot-flatly,boot-journal,boot-lumen,boot-paper,boot-readable,boot-sandstone,boot-slate,boot-spacelab,boot-superhero,boot-yeti,air,modest,retro,splendor";

  function switchCss(css) {
    const stylesheetDom = document.getElementsByTagName("link")[0];
    stylesheetDom.href = "https://dohliam.github.io/dropin-minimal-css/min/" + css + ".min.css";

    saveCss(css)
  }

  function capitalize(s) {
    const u = s.replace(/^(.)/, function(_, l) {
      return l.toUpperCase();
    });

    return u;
  }

  function loadCss() {
    return window.localStorage.getItem('css');
  }

  function saveCss(css) {
    window.localStorage.setItem('css', css);
  }

  function getDefaultCss() {
    const savedCss = loadCss();
    return savedCss ? savedCss : frameworks.split(",")[0];
  }

  function inlineSwitcher() {
    const defaultCss = getDefaultCss();
    const switcher = document.getElementById("switcher");
    const frameworksArr = frameworks.split(",");

    let dropdown = '<select name="switcher_dropdown" id="switcher_dropdown" accesskey="s" onchange="switchCss(this.value)">';
    for (i = 0; i < frameworksArr.length; i++) {
      const framework = frameworksArr[i]
      const frameworkName = capitalize(framework);

      if (framework === defaultCss) {
        option = '<option value="' + framework + '" selected>' + frameworkName + ' CSS</option>';
      } else {
        option = '<option value="' + framework + '">' + frameworkName + ' CSS</option>';
      }
      
      dropdown += option;
    }
    dropdown += '</select>';

    switcher.innerHTML = dropdown;
  }

  function addSwitcher() {
    const defaultCss = getDefaultCss();
    const stylesheetDom = document.getElementsByTagName("link")[0];

    if (stylesheetDom === undefined) {
      const head = document.getElementsByTagName('head')[0];

      const link = document.createElement('link');
      link.rel = "stylesheet";
      link.type = "text/css";
      link.href = "https://dohliam.github.io/dropin-minimal-css/min/" + defaultCss + ".min.css";

      head.appendChild(link);
    }

    const switcher = document.createElement("div");
    switcher.innerHTML = '<div id="switcher">&nbsp;</div>\n<script>inlineSwitcher();<\/script>';

    document.body.insertBefore(switcher, document.body.firstChild);
    document.body.style.paddingLeft = "24px";

    inlineSwitcher();
  }

  addSwitcher();
</script>`
