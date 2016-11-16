module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      build: {
        src: 'build/app.dest.js',
        dest: 'build/app.dest.min.js'
      }
    },
    concat:{
      options: {
        //定义一个字符串插入没个文件之间用于连接输出
        separator: ''
      },
      dist: {
          src: ['src/*.js'],
          dest: 'build/<%= pkg.name %>.cat.js'
      }
    },
    qunit: {
      files: ['test/*.html']
    },
    jshint: {
        files: ['gruntfile.js', 'src/*.js', 'build/*.js'],
        options: {
            globals: {
                exports: true
            }
        }
    },
    /*
    watch: {
        files: ['<%= jshint.files %>'],
        tasks: ['jshint', 'qunit']
    },*/
    
    watch: {
      react: {
        files: 'js/**.*',
        tasks: ['default']
      }
    },
    browserify: {
      options: {
        transform: [ require('grunt-react').browserify ]
      },
      client: {
        src: ['js/app.js','js/commponents/datagrid.js'],
        dest: 'build/app.dest.js'
      }
    },
    copy: {
      main: {
        src: 'build/app.dest.min.js',
        dest: '../assets/js/app.dest.min.js',
      },
    }
  });

  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-concat');
  //grunt.loadNpmTasks('grunt-contrib-qunit');
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-browserify');
  grunt.loadNpmTasks('grunt-contrib-copy');
  // Default task(s).
  //grunt.registerTask('default', ['uglify','concat','qunit','jshint']);
  grunt.registerTask('build', ['browserify']);
  grunt.registerTask('default', ['browserify','uglify',"copy"]);

};