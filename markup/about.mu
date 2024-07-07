div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div mt-6 text-2xl text-center
        div bg-purple-500 rounded-lg p-3
          HomeDucky.com - One Product a Day
        div mt-9 font-mono
          You can
          a href=/core/register link
            create an account
        div mt-9 font-mono
          p
            Then you can add any amazon product you think should be on the homepage.
          p
            The code is open source 
            a href=https://github.com/andrewarrow/homeducky
              https://github.com/andrewarrow/homeducky
